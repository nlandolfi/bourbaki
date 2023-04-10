# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make empirical_cumulative_distribution_functions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

empirical_cumulative_distribution_functions.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > empirical_cumulative_distribution_functions.tex

empirical_cumulative_distribution_functions.pdf: ../../*.tex ../../trademark.pdf *.tex empirical_cumulative_distribution_functions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode empirical_cumulative_distribution_functions.tex
	../../latexrun empirical_cumulative_distribution_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: empirical_cumulative_distribution_functions.pdf
	open empirical_cumulative_distribution_functions.pdf

o:
	make open

terms: empirical_cumulative_distribution_functions.tex
	bbk_sheet -mode ts empirical_cumulative_distribution_functions.tex

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
	rm -f empirical_cumulative_distribution_functions.*
	make reset
	bbk_sheet -mode mk > Makefile