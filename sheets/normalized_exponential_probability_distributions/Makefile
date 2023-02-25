# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make normalized_exponential_probability_distributions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

normalized_exponential_probability_distributions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > normalized_exponential_probability_distributions.tex

normalized_exponential_probability_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex normalized_exponential_probability_distributions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode normalized_exponential_probability_distributions.tex
	../../latexrun normalized_exponential_probability_distributions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: normalized_exponential_probability_distributions.pdf
	open normalized_exponential_probability_distributions.pdf

o:
	make open

terms: normalized_exponential_probability_distributions.tex
	bbk_sheet -mode ts normalized_exponential_probability_distributions.tex

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
	rm -f normalized_exponential_probability_distributions.*
	make reset
	bbk_sheet -mode mk > Makefile