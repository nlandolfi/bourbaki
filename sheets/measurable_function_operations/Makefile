# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make measurable_function_operations.pdf
	make graph.pdf

measurable_function_operations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > measurable_function_operations.tex

measurable_function_operations.pdf: ../../*.tex ../../trademark.pdf *.tex measurable_function_operations.tex
	pdflatex --file-line-error -interaction=nonstopmode measurable_function_operations.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: measurable_function_operations.pdf
	open measurable_function_operations.pdf

o:
	make open

defs: measurable_function_operations.tex
	defs measurable_function_operations.tex

spell:
	aspell -c measurable_function_operations.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f measurable_function_operations.*
	make reset
	bbk_sheet -mode mk > Makefile