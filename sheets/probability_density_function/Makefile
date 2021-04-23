# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make probability_density_function.pdf
	make graph.pdf

probability_density_function.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > probability_density_function.tex

probability_density_function.pdf: ../../*.tex ../../trademark.pdf *.tex probability_density_function.tex
	# pdflatex --file-line-error -interaction=nonstopmode probability_density_function.tex
	../../latexrun probability_density_function.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: probability_density_function.pdf
	open probability_density_function.pdf

o:
	make open

terms: probability_density_function.tex
	bbk_sheet -mode ts probability_density_function.tex

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
	rm -f probability_density_function.*
	make reset
	bbk_sheet -mode mk > Makefile