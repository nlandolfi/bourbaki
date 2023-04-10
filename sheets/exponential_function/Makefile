# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make exponential_function.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

exponential_function.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > exponential_function.tex

exponential_function.pdf: ../../*.tex ../../trademark.pdf *.tex exponential_function.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode exponential_function.tex
	../../latexrun exponential_function.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: exponential_function.pdf
	open exponential_function.pdf

o:
	make open

terms: exponential_function.tex
	bbk sheet -mode ts exponential_function.tex

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
	rm -f exponential_function.*
	make reset
	bbk sheet -mode mk > Makefile