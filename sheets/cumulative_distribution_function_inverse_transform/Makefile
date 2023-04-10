# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make cumulative_distribution_function_inverse_transform.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

cumulative_distribution_function_inverse_transform.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > cumulative_distribution_function_inverse_transform.tex

cumulative_distribution_function_inverse_transform.pdf: ../../*.tex ../../trademark.pdf *.tex cumulative_distribution_function_inverse_transform.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode cumulative_distribution_function_inverse_transform.tex
	../../latexrun cumulative_distribution_function_inverse_transform.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: cumulative_distribution_function_inverse_transform.pdf
	open cumulative_distribution_function_inverse_transform.pdf

o:
	make open

terms: cumulative_distribution_function_inverse_transform.tex
	bbk_sheet -mode ts cumulative_distribution_function_inverse_transform.tex

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
	rm -f cumulative_distribution_function_inverse_transform.*
	make reset
	bbk_sheet -mode mk > Makefile