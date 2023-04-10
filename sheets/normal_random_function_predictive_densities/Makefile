# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make normal_random_function_predictive_densities.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

normal_random_function_predictive_densities.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > normal_random_function_predictive_densities.tex

normal_random_function_predictive_densities.pdf: ../../*.tex ../../trademark.pdf *.tex normal_random_function_predictive_densities.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode normal_random_function_predictive_densities.tex
	../../latexrun normal_random_function_predictive_densities.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: normal_random_function_predictive_densities.pdf
	open normal_random_function_predictive_densities.pdf

o:
	make open

terms: normal_random_function_predictive_densities.tex
	bbk_sheet -mode ts normal_random_function_predictive_densities.tex

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
	rm -f normal_random_function_predictive_densities.*
	make reset
	bbk_sheet -mode mk > Makefile