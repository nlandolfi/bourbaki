# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make moment_generating_function.pdf

moment_generating_function.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > moment_generating_function.tex

moment_generating_function.pdf: ../../*.tex ../../trademark.pdf *.tex moment_generating_function.tex
	# pdflatex --file-line-error -interaction=nonstopmode moment_generating_function.tex
	../../latexrun moment_generating_function.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: moment_generating_function.pdf
	open moment_generating_function.pdf

o:
	make open

terms: moment_generating_function.tex
	bbk_sheet -mode ts moment_generating_function.tex

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
	bbk_sheet -mode mk > Makefile