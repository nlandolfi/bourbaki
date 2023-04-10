# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make cumulative_distribution_function_from_density.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

cumulative_distribution_function_from_density.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > cumulative_distribution_function_from_density.tex

cumulative_distribution_function_from_density.pdf: ../../*.tex ../../trademark.pdf *.tex cumulative_distribution_function_from_density.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode cumulative_distribution_function_from_density.tex
	../../latexrun cumulative_distribution_function_from_density.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: cumulative_distribution_function_from_density.pdf
	open cumulative_distribution_function_from_density.pdf

o:
	make open

terms: cumulative_distribution_function_from_density.tex
	bbk sheet -mode ts cumulative_distribution_function_from_density.tex

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
	rm -f cumulative_distribution_function_from_density.*
	make reset
	bbk sheet -mode mk > Makefile