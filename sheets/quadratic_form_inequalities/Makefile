# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make quadratic_form_inequalities.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

quadratic_form_inequalities.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > quadratic_form_inequalities.tex

quadratic_form_inequalities.pdf: ../../*.tex ../../trademark.pdf *.tex quadratic_form_inequalities.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode quadratic_form_inequalities.tex
	../../latexrun quadratic_form_inequalities.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: quadratic_form_inequalities.pdf
	open quadratic_form_inequalities.pdf

o:
	make open

terms: quadratic_form_inequalities.tex
	bbk sheet -mode ts quadratic_form_inequalities.tex

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
	rm -f quadratic_form_inequalities.*
	make reset
	bbk sheet -mode mk > Makefile