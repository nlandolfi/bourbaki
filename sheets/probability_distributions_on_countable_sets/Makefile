# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make probability_distributions_on_countable_sets.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

probability_distributions_on_countable_sets.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > probability_distributions_on_countable_sets.tex

probability_distributions_on_countable_sets.pdf: ../../*.tex ../../trademark.pdf *.tex probability_distributions_on_countable_sets.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode probability_distributions_on_countable_sets.tex
	../../latexrun probability_distributions_on_countable_sets.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: probability_distributions_on_countable_sets.pdf
	open probability_distributions_on_countable_sets.pdf

o:
	make open

terms: probability_distributions_on_countable_sets.tex
	bbk sheet -mode ts probability_distributions_on_countable_sets.tex

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
	rm -f probability_distributions_on_countable_sets.*
	make reset
	bbk sheet -mode mk > Makefile