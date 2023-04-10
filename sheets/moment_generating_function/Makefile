# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make moment_generating_function.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

moment_generating_function.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > moment_generating_function.tex

moment_generating_function.pdf: ../../*.tex ../../trademark.pdf *.tex moment_generating_function.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode moment_generating_function.tex
	../../latexrun moment_generating_function.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: moment_generating_function.pdf
	open moment_generating_function.pdf

o:
	make open

terms: moment_generating_function.tex
	bbk sheet -mode ts moment_generating_function.tex

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
	rm -f moment_generating_function.*
	make reset
	bbk sheet -mode mk > Makefile