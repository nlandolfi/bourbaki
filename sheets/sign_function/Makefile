# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make sign_function.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

sign_function.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > sign_function.tex

sign_function.pdf: ../../*.tex ../../trademark.pdf *.tex sign_function.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode sign_function.tex
	../../latexrun sign_function.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: sign_function.pdf
	open sign_function.pdf

o:
	make open

terms: sign_function.tex
	bbk sheet -mode ts sign_function.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f sign_function.*
	make reset
	bbk sheet -mode mk > Makefile