# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make cross_entropy_of_probability_distributions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

cross_entropy_of_probability_distributions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > cross_entropy_of_probability_distributions.tex

cross_entropy_of_probability_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex cross_entropy_of_probability_distributions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode cross_entropy_of_probability_distributions.tex
	../../latexrun cross_entropy_of_probability_distributions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: cross_entropy_of_probability_distributions.pdf
	open cross_entropy_of_probability_distributions.pdf

o:
	make open

terms: cross_entropy_of_probability_distributions.tex
	bbk_sheet -mode ts cross_entropy_of_probability_distributions.tex

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
	rm -f cross_entropy_of_probability_distributions.*
	make reset
	bbk_sheet -mode mk > Makefile